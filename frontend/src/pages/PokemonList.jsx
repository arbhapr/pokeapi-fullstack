import { useEffect, useState } from "react";
import { Card } from "../components";

const PokemonList = () => {
    const [pokemonData, setPokemonData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [nextPageUrl, setNextPageUrl] = useState(null);
    const [prevPageUrl, setPrevPageUrl] = useState(null);

    const fetchPokemonData = async (url) => {
        try {
            setLoading(true);
            const response = await fetch(url);
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            setPokemonData(data.data || []); // Ensure pokemonData is always an array
            setNextPageUrl(data.pagination.nextPage);
            setPrevPageUrl(data.pagination.prevPage);
        } catch (error) {
            setError(error.message);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchPokemonData(`${import.meta.env.VITE_API_URL}/pokemon`);
    }, []);

    const handleNextPage = () => {
        if (nextPageUrl) {
            fetchPokemonData(nextPageUrl);
        }
    };

    const handlePrevPage = () => {
        if (prevPageUrl) {
            fetchPokemonData(prevPageUrl);
        }
    };

    if (loading) {
        return <div className="text-center py-8">Loading...</div>;
    }

    if (error) {
        return (
            <div className="text-center py-8 text-red-500">Error: {error}</div>
        );
    }

    return (
        <div className="container mx-auto py-8">
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                {pokemonData.length > 0 ? (
                    pokemonData.map((pokemon) => (
                        <Card
                            key={pokemon.id}
                            name={pokemon.name}
                            sprite={pokemon.sprite}
                            url={pokemon.url}
                            id={pokemon.id}
                        />
                    ))
                ) : (
                    <div className="text-center col-span-full py-8">
                        No Pokémon available
                    </div>
                )}
            </div>
            <div className="flex justify-between mt-8">
                <button
                    onClick={handlePrevPage}
                    disabled={!prevPageUrl}
                    className={`py-2 px-4 rounded ${
                        !prevPageUrl
                            ? "bg-gray-400 text-gray-700 cursor-not-allowed"
                            : "bg-blue-500 text-white"
                    }`}
                >
                    Previous
                </button>
                <button
                    onClick={handleNextPage}
                    disabled={!nextPageUrl}
                    className={`py-2 px-4 rounded ${
                        !nextPageUrl
                            ? "bg-gray-400 text-gray-700 cursor-not-allowed"
                            : "bg-blue-500 text-white"
                    }`}
                >
                    Next
                </button>
            </div>
        </div>
    );
};

export default PokemonList;
