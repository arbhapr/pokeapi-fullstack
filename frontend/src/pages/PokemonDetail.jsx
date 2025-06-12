import { useParams, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

// Reusable Button Component
const Button = ({ onClick, disabled, loading, children, className }) => (
    <button
        onClick={onClick}
        disabled={disabled}
        className={`py-2 px-4 rounded ${className}`}
    >
        {loading ? "Loading..." : children}
    </button>
);

const PokemonDetail = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [pokemon, setPokemon] = useState(null);
    const [isLoading, setIsLoading] = useState(false);
    const [isCatched, setIsCatched] = useState(null);
    const [fetchError, setFetchError] = useState(null);

    useEffect(() => {
        const fetchPokemonDetail = async () => {
            try {
                const response = await fetch(
                    `${import.meta.env.VITE_API_URL}/pokemon/${id}`
                );
                if (!response.ok) throw new Error("Failed to fetch Pokémon.");
                const data = await response.json();
                setPokemon(data);
            } catch (error) {
                setFetchError(error.message);
                console.error("Failed to fetch Pokémon details:", error);
            }
        };

        fetchPokemonDetail();
    }, [id]);

    const handleBack = () => navigate(-1);

    const handleCatch = async () => {
        setIsLoading(true);
        setIsCatched(null);

        try {
            const response = await fetch(
                `${import.meta.env.VITE_API_URL}/pokemon/${id}`,
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                }
            );

            const data = await response.json();
            setIsCatched(data.success ? true : false);
        } catch (error) {
            console.error("Failed to catch Pokémon:", error);
            setIsCatched(false);
        } finally {
            setIsLoading(false);
        }
    };

    if (fetchError) return <div className="text-red-500">{fetchError}</div>;

    if (!pokemon) return <div>Loading...</div>;

    return (
        <div className="min-h-screen bg-gray-100 p-8">
            <div className="container mx-auto max-w-lg bg-white shadow-lg rounded-lg p-6">
                <div className="flex flex-col items-center">
                    <img
                        src={pokemon.sprite}
                        alt={pokemon.name}
                        className="w-48 h-48 object-cover rounded-full mb-4"
                    />
                    <h1 className="text-3xl font-bold mb-2">{pokemon.name}</h1>

                    {pokemon.types?.length > 0 && (
                        <>
                            <div className="text-lg font-semibold mb-4">
                                Types
                            </div>
                            <div className="flex flex-wrap gap-2 mb-6">
                                {pokemon.types.map((type) => (
                                    <span
                                        key={type.name}
                                        className="bg-gray-200 text-gray-800 py-1 px-3 rounded-full text-sm"
                                    >
                                        {type.name}
                                    </span>
                                ))}
                            </div>
                        </>
                    )}

                    {pokemon.moves?.length > 0 && (
                        <>
                            <div className="text-lg font-semibold mb-4">
                                Moves
                            </div>
                            <div className="flex flex-wrap gap-2 mb-6">
                                {pokemon.moves.map((move) => (
                                    <span
                                        key={move.name}
                                        className="bg-blue-100 text-blue-800 py-1 px-3 rounded-full text-sm"
                                    >
                                        {move.name}
                                    </span>
                                ))}
                            </div>
                        </>
                    )}

                    <div className="flex justify-between w-full mt-6">
                        <Button
                            onClick={handleBack}
                            className="bg-gray-500 text-white"
                        >
                            Back
                        </Button>
                        <Button
                            onClick={handleCatch}
                            loading={isLoading}
                            disabled={isLoading}
                            className="bg-green-500 text-white"
                        >
                            Catch
                        </Button>
                    </div>

                    {/* Modal for Catch Result */}
                    {isCatched !== null && (
                        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
                            <div className="bg-white p-8 rounded-lg shadow-lg text-center">
                                {isCatched ? (
                                    <>
                                        <h2 className="text-2xl font-bold mb-4">
                                            Successfully caught Pokémon!
                                        </h2>
                                        <img
                                            src={pokemon.sprite}
                                            alt={pokemon.name}
                                            className="w-24 h-24 mx-auto mb-4"
                                        />
                                        <p className="text-lg">
                                            {pokemon.name}
                                        </p>
                                    </>
                                ) : (
                                    <>
                                        <h2 className="text-2xl font-bold mb-4 text-red-500">
                                            Failed to catch the Pokémon.
                                        </h2>
                                    </>
                                )}
                                <Button
                                    onClick={() => setIsCatched(null)}
                                    className="bg-blue-500 text-white mt-4"
                                >
                                    Close
                                </Button>
                            </div>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export default PokemonDetail;
