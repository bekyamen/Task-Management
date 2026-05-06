import { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { api } from '../api/client';
import { LogIn } from 'lucide-react';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await api.post('/auth/login', { email, password });
      login(res.data.token);
      navigate('/');
    } catch (err: any) {
      setError(err.response?.data?.error || 'Login failed');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-background">
      <div className="max-w-md w-full space-y-8 glass-card p-8">
        <div>
          <div className="flex justify-center text-primary">
            <LogIn size={48} />
          </div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-textMain">Sign in to your account</h2>
        </div>
        <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
          {error && <div className="text-red-500 text-sm text-center">{error}</div>}
          <div className="rounded-md shadow-sm space-y-4">
            <div>
              <input
                type="email"
                required
                className="input-field"
                placeholder="Email address"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
              />
            </div>
            <div>
              <input
                type="password"
                required
                className="input-field"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            <div className="flex justify-end">
              <Link to="/forgot-password" className="text-sm font-medium text-primary hover:text-primaryHover">
                Forgot your password?
              </Link>
            </div>
          </div>

          <div>
            <button type="submit" className="w-full btn-primary flex justify-center py-2 px-4">
              Sign in
            </button>
          </div>
          <div className="text-center text-sm">
            <span className="text-textMuted">Don't have an account? </span>
            <Link to="/register" className="font-medium text-primary hover:text-primaryHover">
              Register here
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
}
