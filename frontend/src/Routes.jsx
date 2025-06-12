import { Route, Routes } from "react-router-dom";
import { PokemonDetail, MyBackpack, PokemonList } from "./pages";

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<PokemonList />} />
            <Route path="/pokemon-list" element={<PokemonList />} />
            <Route path="/pokemon/:id" element={<PokemonDetail />} />
            <Route path="/my-backpack" element={<MyBackpack />} />
        </Routes>
    );
};

export default AppRoutes;
