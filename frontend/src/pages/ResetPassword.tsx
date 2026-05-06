import { useState } from 'react';
import { useNavigate, useSearchParams, Link } from 'react-router-dom';
import { api } from '../api/client';
import { Lock, CheckCircle, ArrowLeft } from 'lucide-react';

export default function ResetPassword() {
  const [searchParams] = useSearchParams();
  const token = searchParams.get('token');
  
  const [newPassword, setNewPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [status, setStatus] = useState<'idle' | 'loading' | 'success' | 'error'>('idle');
  const [error, setError] = useState('');
  
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!token) {
      setStatus('error');
      setError('Invalid or missing reset token.');
      return;
    }
    
    if (newPassword !== confirmPassword) {
      setStatus('error');
      setError('Passwords do not match.');
      return;
    }
    
    if (newPassword.length < 6) {
      setStatus('error');
      setError('Password must be at least 6 characters.');
      return;
    }

    setStatus('loading');
    setError('');
    
    try {
      await api.post('/auth/reset-password', { token, newPassword });
      setStatus('success');
    } catch (err: any) {
      setStatus('error');
      setError(err.response?.data?.error || 'Failed to reset password');
    }
  };

  if (!token && status === 'idle') {
    return (
      <div className="min-h-screen flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-background">
        <div className="max-w-md w-full space-y-8 glass-card p-8 text-center">
           <div className="flex justify-center text-red-500 mb-4">
             <Lock size={48} />
           </div>
           <h2 className="text-2xl font-bold text-textMain">Invalid Link</h2>
           <p className="text-textMuted mt-2">The password reset link is invalid or missing the token.</p>
           <div className="mt-8">
             <Link to="/forgot-password" className="btn-primary inline-block py-2 px-6">
                Request New Link
             </Link>
           </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-background">
      <div className="max-w-md w-full space-y-8 glass-card p-8">
        <div>
          <div className="flex justify-center text-primary">
            {status === 'success' ? <CheckCircle size={48} className="text-green-500" /> : <Lock size={48} />}
          </div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-textMain">
            {status === 'success' ? 'Password Reset Complete' : 'Create New Password'}
          </h2>
        </div>
        
        {status === 'success' ? (
          <div className="mt-8 space-y-6">
            <p className="text-center text-textMuted">
              Your password has been successfully reset. You can now use your new password to log in.
            </p>
            <button 
              onClick={() => navigate('/login')}
              className="w-full btn-primary flex justify-center py-2 px-4"
            >
              Proceed to Login
            </button>
          </div>
        ) : (
          <form className="mt-8 space-y-6" onSubmit={handleSubmit}>
            {error && <div className="text-red-500 text-sm text-center bg-red-500/10 p-3 rounded-md border border-red-500/20">{error}</div>}
            
            <div className="rounded-md shadow-sm space-y-4">
              <div>
                <input
                  type="password"
                  required
                  className="input-field"
                  placeholder="New Password (min 6 chars)"
                  value={newPassword}
                  onChange={(e) => setNewPassword(e.target.value)}
                  disabled={status === 'loading'}
                  minLength={6}
                />
              </div>
              <div>
                <input
                  type="password"
                  required
                  className="input-field"
                  placeholder="Confirm New Password"
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)}
                  disabled={status === 'loading'}
                  minLength={6}
                />
              </div>
            </div>

            <div>
              <button 
                type="submit" 
                className="w-full btn-primary flex justify-center py-2 px-4 disabled:opacity-50 disabled:cursor-not-allowed"
                disabled={status === 'loading'}
              >
                {status === 'loading' ? 'Resetting...' : 'Reset Password'}
              </button>
            </div>
            
            <div className="text-center text-sm">
              <Link to="/login" className="flex items-center justify-center font-medium text-textMuted hover:text-textMain transition-colors">
                <ArrowLeft size={16} className="mr-2" />
                Cancel and return to login
              </Link>
            </div>
          </form>
        )}
      </div>
    </div>
  );
}
