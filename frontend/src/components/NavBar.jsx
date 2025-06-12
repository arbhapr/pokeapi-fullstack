import { Link } from "react-router-dom";

const NavBar = () => {
    return (
        <nav className="bg-gray-800 p-4">
            <div className="container mx-auto flex justify-between items-center">
                <div className="text-white font-bold text-xl">
                    <img
                        src="/pokemon-logo.svg"
                        alt="Logo"
                        className="h-9 inline-block mr-2"
                    />
                </div>
                <div className="flex space-x-4">
                    <Link
                        to="/pokemon-list"
                        className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-700 transition duration-300"
                    >
                        Pokemon List
                    </Link>
                    <Link
                        to="/my-backpack"
                        className="bg-green-500 text-white py-2 px-4 rounded hover:bg-green-700 transition duration-300"
                    >
                        My Backpack
                    </Link>
                </div>
            </div>
        </nav>
    );
};

export default NavBar;
