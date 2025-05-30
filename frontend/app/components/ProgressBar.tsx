"use client";

import { useEffect, useState } from "react";

type Props = {
  duration: number; // アニメーションにかける秒数（バーが100%に到達するまでの時間）
  color: string; // バーの色（Tailwind CSSのクラス名）
};

export default function ProgressBar({ duration, color }: Props) {
  // 現在の進捗率（0〜100）
  const [percent, setPercent] = useState(0);
  // バーの表示トリガー（durationが有効なときだけ描画）
  const [showBar, setShowBar] = useState(false);
  // アニメーション完了フラグ（100%に到達したかどうか）
  const [completed, setCompleted] = useState(false);

  useEffect(() => {
    // 初期化（タスク変更時などに毎回リセット）
    setPercent(0);
    setShowBar(false);
    setCompleted(false);

    // 無効なdurationなら処理せず終了
    if (!duration || duration <= 0) return;

    // バーの表示開始
    setShowBar(true);

    // 100ミリ秒ごとに進捗を更新（durationに応じて割る）
    const interval = setInterval(() => {
      setPercent((prev) => {
        const next = prev + 100 / (duration * 10); // 10回/秒の更新
        if (next >= 100) {
          clearInterval(interval); // アニメーション完了で停止
          setCompleted(true); // 秒数表示トリガーON
          return 100;
        }
        return next;
      });
    }, 100);

    // コンポーネントのクリーンアップ
    return () => clearInterval(interval);
  }, [duration]);

  return (
    <div className="w-full flex items-center space-x-4">
      {/* バー全体枠 */}
      <div className="flex-grow bg-gray-200 rounded h-5 overflow-hidden">
        {/* バー本体（progress%だけ塗る） */}
        {showBar && <div className={`${color} h-full transition-all duration-100`} style={{ width: `${percent}%` }} />}
      </div>

      {/* 完了後にだけ秒数表示（グラフの外に固定） */}
      <div className="text-sm text-gray-700 whitespace-nowrap min-w-[60px]">{completed ? `${duration.toFixed(2)} 秒` : ""}</div>
    </div>
  );
}
