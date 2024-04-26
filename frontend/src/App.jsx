import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import Home from './pages/Home';
import Loader from './pages/Loader';
import NotFound404 from './pages/NotFound404';

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Loader />} />
        <Route path="/app" element={<Home />} />
        <Route path="/404" element={<NotFound404 />} />
      </Routes>
    </Router>
  )
}

export default App
