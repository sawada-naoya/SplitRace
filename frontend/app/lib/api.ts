export type TaskProgress = {
  elapsed: number;
  progress: number[];
};

export type RunTaskResponse = {
  serial: TaskProgress;
  parallel: TaskProgress;
};

export async function fetchTaskProgress(count: number): Promise<RunTaskResponse> {
  const res = await fetch("http://localhost:8080/run-dual-tasks", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ count }), // Goのdto.RunTasckRequestに対応
  });
  if (!res.ok) throw new Error("API取得失敗");
  return res.json();
}
