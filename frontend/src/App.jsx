import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

import Home from './pages/Home';
import Loader from './pages/Loader';
import LinkHandler from './pages/LinkHandler';

function App() {

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Loader />} />
        <Route path="/app" element={<Home />} />
        <Route path="/:linkId" element={<LinkHandler />} />
      </Routes>
    </Router>
  )
}

export default App
