import { BrowserRouter as Router } from "react-router-dom";
import { NavBar } from "./components";
import AppRoutes from "./Routes";

function App() {
    return (
        <Router>
            <div className="min-h-screen bg-gray-100">
                <NavBar />
                <div className="container mx-auto py-8">
                    <AppRoutes />
                </div>
            </div>
        </Router>
    );
}

export default App;
