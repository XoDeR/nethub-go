import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Home from "./pages/home";
import Login from "./pages/login";
import Register from "./pages/register";
import Feed from "./pages/feed";

// TODO set VITE_API_URL
export const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080/v1"

const isAuthenticated = (): boolean => {
  return localStorage.getItem("token") !== null;
};

const ProtectedRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return isAuthenticated() ? children : <Navigate to="/login" />;
};

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        {/* Public routes */}
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />

        {/* Protected route */}
        <Route path="/feed" element={<ProtectedRoute><Feed /></ProtectedRoute>} />
      </Routes>
    </Router>
  );
};

export default App;

