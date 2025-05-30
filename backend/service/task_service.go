package service

import (
	"net/http"
	"sync"
	"time"

	"github.com/sawada-naoya/splitrace/dto"
)

// TaskService インターフェース：直列/並列処理を提供する
type TaskService interface {
	RunSerial(count int) dto.TaskProgress
	RunParallel(count int) dto.TaskProgress
}

// taskServiceImpl 実体構造体（DIで注入される）
type taskServiceImpl struct{}

// NewTaskService DI用のコンストラクタ
func NewTaskService() TaskService {
	return &taskServiceImpl{}
}

// 直列処理: 外部APIを直列でcount回実行し、経過秒数を蓄積して返す
func (s *taskServiceImpl) RunSerial(count int) dto.TaskProgress {
	start := time.Now() // 計測開始時間
	var res dto.TaskProgress
	for i := 0; i < count; i++ {
		// 外部ダミーAPIを1回叩く（ブロッキング＝前が終わるまで次に進まない）
		http.Get("https://httpbin.org/delay/1")
		// この時点の経過時間）を記録
		res.Progress = append(res.Progress, time.Since(start).Seconds())
	}
	// 総経過時間
	res.Elapsed = time.Since(start).Seconds()
	return res
}

// 並列処理: 外部APIを並列でcount回同時実行し、完了タイミングを記録して返す
// ① goroutineで並列実行 → 値をチャネルに送信
// ② 全部終わるまで wait（ch はまだ読まない）
// ③ wait完了後、ch を閉じる
// ④ ch を range で1つずつ読み取り、結果を収集
func (s *taskServiceImpl) RunParallel(count int) dto.TaskProgress {
	start := time.Now()
	var res dto.TaskProgress
	var wg sync.WaitGroup
	// 各タスクの完了秒数を非同期で集めるチャンネル
	ch := make(chan float64, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		// 並列実行開始（goroutine）
		go func() {
			defer wg.Done()
			http.Get("https://httpbin.org/delay/1")
			// 完了時点の経過秒数をチャネルに送信
			ch <- time.Since(start).Seconds()
		}()
	}
	// すべての並列処理が終わるのを待つ
	wg.Wait()
	// チャンネルを閉じて for-range が終わるようにする
	close(ch)
	// 全リクエストの完了時刻を収集
	for p := range ch {
		res.Progress = append(res.Progress, p)
	}
	res.Elapsed = time.Since(start).Seconds()
	return res
}
