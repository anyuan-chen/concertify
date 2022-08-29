import { BrowserRouter, Routes, Route } from "react-router-dom";
import Layout from "./components/layout";
import Home from "./pages/index";
import Select from "./pages/select";

function App() {
  return (
    <Layout>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home></Home>}></Route>
          <Route path="/select" element={<Select></Select>}></Route>
        </Routes>
      </BrowserRouter>
    </Layout>
  );
}

export default App;
