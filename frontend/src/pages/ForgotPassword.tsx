import { useState } from 'react';
import { Link } from 'react-router-dom';
import { api } from '../api/client';
import { Mail, ArrowLeft } from 'lucide-react';

export default function ForgotPassword() {
  const [email, setEmail] = useState('');
  const [status, setStatus] = useState<'idle' | 'loading' | 'success' | 'error'>('idle');
  const [error, setError] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setStatus('loading');
    setError('');
    setMessage('');
    
    try {
      const res = await api.post('/auth/forgot-password', { email });
      setStatus('success');
      setMessage(res.data.message);
    } catch (err: any) {
      setStatus('error');
      setError(err.response?.data?.error || 'Failed to process request');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-background">
      <div className="max-w-md w-full space-y-8 glass-card p-8">
        <div>
          <div className="flex justify-center text-primary">
            <Mail size={48} />
          </div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-textMain">Reset Password</h2>
          <p className="mt-2 text-center text-sm text-textMuted">
            Enter your email address to receive a password reset link.
          </p>
        </div>
        
        {status === 'success' ? (
          <div className="rounded-md bg-green-500/10 p-4 border border-green-500/20">
            <p className="text-sm text-green-400 text-center">{message}</p>
            <div className="mt-6 flex justify-center">
              <Link to="/login" className="flex items-center text-primary hover:text-primaryHover text-sm font-medium">
                <ArrowLeft size={16} className="mr-2" />
                Return to Login
              </Link>
            </div>
          </div>
        ) : (
          <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
            {error && <div className="text-red-500 text-sm text-center bg-red-500/10 p-3 rounded-md border border-red-500/20">{error}</div>}
            <div className="rounded-md shadow-sm space-y-4">
              <div>
                <input
                  type="email"
                  required
                  className="input-field"
                  placeholder="Email address"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  disabled={status === 'loading'}
                />
              </div>
            </div>

            <div>
              <button 
                type="submit" 
                className="w-full btn-primary flex justify-center py-2 px-4 disabled:opacity-50 disabled:cursor-not-allowed"
                disabled={status === 'loading'}
              >
                {status === 'loading' ? 'Sending...' : 'Send Reset Link'}
              </button>
            </div>
            
            <div className="text-center text-sm">
              <Link to="/login" className="flex items-center justify-center font-medium text-primary hover:text-primaryHover">
                <ArrowLeft size={16} className="mr-2" />
                Back to login
              </Link>
            </div>
          </form>
        )}
      </div>
    </div>
  );
}
