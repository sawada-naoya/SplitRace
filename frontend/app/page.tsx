"use client";

import ProgressBar from "./components/ProgressBar";
import { useForm } from "react-hook-form";
import { useState } from "react";
import { fetchTaskProgress, RunTaskResponse } from "./lib/api";

type FormData = {
  taskCount: number;
};

export default function HomePage() {
  const { register, handleSubmit, reset } = useForm<FormData>();
  const [chartData, setChartData] = useState<RunTaskResponse | null>(null);

  const onSubmit = async (data: FormData) => {
    try {
      const json = await fetchTaskProgress(data.taskCount); // API呼び出し
      setChartData(json); // 結果をステートに保存
    } catch (err) {
      console.error(err);
      alert("API呼び出しに失敗しました");
    }
  };

  const handleReset = () => {
    reset(); // フォームの初期化
    setChartData(null); // グラフのリセット
  };

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col items-center justify-start py-10 px-4">
      <h1 className="text-3xl font-bold mb-6 text-gray-800">並列 vs 直列 処理速度比較</h1>

      <form onSubmit={handleSubmit(onSubmit)} className="flex items-center space-x-4 mb-8">
        <input type="number" placeholder="タスク数" {...register("taskCount", { required: true, min: 1, valueAsNumber: true })} className="px-4 py-2 border border-gray-300 rounded-md shadow-sm" />
        <button type="submit" className="bg-blue-600 text-white px-6 py-2 rounded-md hover:bg-blue-700 transition">
          実行
        </button>
        <button type="button" onClick={handleReset} className="bg-gray-400 text-white px-6 py-2 rounded-md hover:bg-gray-500 transition">
          リセット
        </button>
      </form>

      <div className="w-full max-w-2xl space-y-8">
        <div>
          <h2 className="text-lg font-semibold mb-2 text-gray-700 flex justify-between">
            <span>直列処理</span>
          </h2>
          <ProgressBar duration={chartData?.serial.elapsed ?? 0} color="bg-blue-500" />
        </div>

        <div>
          <h2 className="text-lg font-semibold mb-2 text-gray-700 flex justify-between">
            <span>並列処理</span>
          </h2>
          <ProgressBar duration={chartData?.parallel.elapsed ?? 0} color="bg-red-500" />
        </div>
      </div>
    </div>
  );
}
