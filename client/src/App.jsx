import { useEffect, useState } from "react";

const API_URL = import.meta.env.VITE_API_URL;
const API_KEY = import.meta.env.VITE_API_KEY;

console.log("API_URL:", API_URL);
console.log("API_KEY:", API_KEY);

export default function App() {
  const [tasks, setTasks] = useState([]);

  const [form, setForm] = useState({
    title: "",
    description: "",
    status: "todo",
    due_date: "",
  });

  async function fetchTasks() {
    try {
      const response = await fetch(API_URL, {
        headers: {
          Authorization: `Bearer ${API_KEY}`,
        },
      });
      const data = await response.json();

      setTasks(data.data || []);
    } catch (error) {
      console.log(error);
    }
  }

  useEffect(() => {
    fetchTasks();
  }, []);

  async function handleSubmit(e) {
    e.preventDefault();

    try {
      await fetch(API_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${API_KEY}`,
        },
        body: JSON.stringify({
          ...form,
          due_date: new Date(form.due_date).toISOString(),
        }),
      });

      setForm({
        title: "",
        description: "",
        status: "todo",
        due_date: "",
      });

      fetchTasks();
    } catch (error) {
      console.log(error);
    }
  }

  async function deleteTask(id) {
    try {
      await fetch(
        `${API_URL}/${id}`,

        {
          headers: {
            Authorization: `Bearer ${API_KEY}`,
          },
          method: "DELETE",
        },
      );

      fetchTasks();
    } catch (error) {
      console.log(error);
    }
  }

  async function updateStatus(task, status) {
    try {
      await fetch(`${API_URL}/${task.id}`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${API_KEY}`,
        },
        body: JSON.stringify({
          title: task.title,
          description: task.description,
          status,
          due_date: task.due_date,
        }),
      });

      fetchTasks();
    } catch (error) {
      console.log(error);
    }
  }

  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <div className="max-w-5xl mx-auto">
        <h1 className="text-4xl font-bold mb-8 text-center">Task Management</h1>

        <form
          onSubmit={handleSubmit}
          className="bg-white p-6 rounded-2xl shadow-md mb-8"
        >
          <div className="grid md:grid-cols-2 gap-4">
            <input
              type="text"
              placeholder="Title"
              className="border rounded-xl p-3 outline-none"
              value={form.title}
              onChange={(e) => setForm({ ...form, title: e.target.value })}
            />

            <input
              type="date"
              className="border rounded-xl p-3 outline-none"
              value={form.due_date}
              onChange={(e) => setForm({ ...form, due_date: e.target.value })}
            />

            <textarea
              placeholder="Description"
              className="border rounded-xl p-3 outline-none md:col-span-2"
              rows="4"
              value={form.description}
              onChange={(e) =>
                setForm({
                  ...form,
                  description: e.target.value,
                })
              }
            />

            <select
              className="border rounded-xl p-3 outline-none"
              value={form.status}
              onChange={(e) => setForm({ ...form, status: e.target.value })}
            >
              <option value="todo">Todo</option>
              <option value="inprogress">In Progress</option>
              <option value="done">Done</option>
            </select>
          </div>

          <button
            type="submit"
            className="mt-4 bg-black text-white px-6 py-3 rounded-xl hover:opacity-90"
          >
            Add Task
          </button>
        </form>

        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
          {tasks.map((task) => (
            <div key={task.id} className="bg-white rounded-2xl shadow-md p-5">
              <div className="flex justify-between items-start mb-4">
                <h2 className="text-xl font-semibold">{task.title}</h2>

                <span
                  className={`text-sm px-3 py-1 rounded-full ${
                    task.status === "todo"
                      ? "bg-yellow-100 text-yellow-700"
                      : task.status === "inprogress"
                        ? "bg-blue-100 text-blue-700"
                        : "bg-green-100 text-green-700"
                  }`}
                >
                  {task.status}
                </span>
              </div>

              <p className="text-gray-600 mb-4">{task.description}</p>

              <p className="text-sm text-gray-500 mb-4">
                Due: {new Date(task.due_date).toLocaleDateString()}
              </p>

              <div className="flex gap-2 flex-wrap">
                <button
                  onClick={() => updateStatus(task, "todo")}
                  className="px-3 py-2 rounded-lg bg-yellow-500 text-white text-sm"
                >
                  Todo
                </button>

                <button
                  onClick={() => updateStatus(task, "inprogress")}
                  className="px-3 py-2 rounded-lg bg-blue-500 text-white text-sm"
                >
                  Progress
                </button>

                <button
                  onClick={() => updateStatus(task, "done")}
                  className="px-3 py-2 rounded-lg bg-green-500 text-white text-sm"
                >
                  Done
                </button>

                <button
                  onClick={() => deleteTask(task.id)}
                  className="px-3 py-2 rounded-lg bg-red-500 text-white text-sm"
                >
                  Delete
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
