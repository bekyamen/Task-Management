import { useState, useEffect } from 'react';
import { useAuth } from '../context/AuthContext';
import { api } from '../api/client';
import { LogOut, Plus, Trash2, Edit2, CheckCircle, Circle } from 'lucide-react';

interface Task {
  id: number;
  title: string;
  description: string;
  status: string;
}

export default function Dashboard() {
  const { logout } = useAuth();
  const [tasks, setTasks] = useState<Task[]>([]);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [editingId, setEditingId] = useState<number | null>(null);

  const fetchTasks = async () => {
    try {
      const res = await api.get('/tasks');
      setTasks(res.data || []);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title) return;
    
    try {
      if (editingId) {
        await api.put(`/tasks/${editingId}`, { title, description, status: 'pending' });
        setEditingId(null);
      } else {
        await api.post('/tasks', { title, description });
      }
      setTitle('');
      setDescription('');
      fetchTasks();
    } catch (err) {
      console.error(err);
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await api.delete(`/tasks/${id}`);
      fetchTasks();
    } catch (err) {
      console.error(err);
    }
  };

  const toggleStatus = async (task: Task) => {
    try {
      const newStatus = task.status === 'completed' ? 'pending' : 'completed';
      await api.put(`/tasks/${task.id}`, { ...task, status: newStatus });
      fetchTasks();
    } catch (err) {
      console.error(err);
    }
  };

  const editTask = (task: Task) => {
    setEditingId(task.id);
    setTitle(task.title);
    setDescription(task.description);
  };

  return (
    <div className="min-h-screen bg-background text-textMain p-4 md:p-8">
      <div className="max-w-4xl mx-auto space-y-8">
        <header className="flex justify-between items-center bg-surface/50 p-6 rounded-xl border border-gray-700/30 shadow-lg">
          <h1 className="text-3xl font-bold text-white">Tasks Management</h1>
          <button onClick={logout} className="flex items-center space-x-2 text-red-400 hover:text-red-300 transition">
            <LogOut size={20} />
            <span>Logout</span>
          </button>
        </header>

        <form onSubmit={handleSubmit} className="glass-card p-6 space-y-4">
          <h2 className="text-xl font-semibold mb-2">{editingId ? 'Edit Task' : 'Create New Task'}</h2>
          <div>
            <input
              type="text"
              placeholder="Task Title"
              className="input-field"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              required
            />
          </div>
          <div>
            <textarea
              placeholder="Description (optional)"
              className="input-field"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              rows={3}
            />
          </div>
          <button type="submit" className="btn-primary flex items-center space-x-2">
            <Plus size={20} />
            <span>{editingId ? 'Update Task' : 'Add Task'}</span>
          </button>
          {editingId && (
            <button type="button" onClick={() => { setEditingId(null); setTitle(''); setDescription(''); }} className="ml-4 text-gray-400 hover:text-white">
              Cancel
            </button>
          )}
        </form>

        <div className="space-y-4">
          {tasks.length === 0 ? (
            <div className="text-center text-textMuted py-8">No tasks found. Create one above!</div>
          ) : (
            tasks.map(task => (
              <div key={task.id} className="glass-card p-5 flex justify-between items-center group transition-all hover:bg-surface/70">
                <div className="flex items-start space-x-4">
                  <button onClick={() => toggleStatus(task)} className="mt-1 flex-shrink-0 text-primary hover:text-primaryHover transition">
                    {task.status === 'completed' ? <CheckCircle className="text-green-500" /> : <Circle />}
                  </button>
                  <div>
                    <h3 className={`text-lg font-semibold ${task.status === 'completed' ? 'line-through text-textMuted' : 'text-white'}`}>
                      {task.title}
                    </h3>
                    {task.description && (
                      <p className={`text-sm mt-1 ${task.status === 'completed' ? 'text-gray-600' : 'text-gray-400'}`}>
                        {task.description}
                      </p>
                    )}
                  </div>
                </div>
                <div className="flex space-x-3 opacity-0 group-hover:opacity-100 transition-opacity">
                  <button onClick={() => editTask(task)} className="text-blue-400 hover:text-blue-300">
                    <Edit2 size={18} />
                  </button>
                  <button onClick={() => handleDelete(task.id)} className="text-red-400 hover:text-red-300">
                    <Trash2 size={18} />
                  </button>
                </div>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
}
