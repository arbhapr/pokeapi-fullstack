import { useEffect, useState } from "react";
import { Card, Spinner, Modal } from "../components"; // Import Modal component

const MyBackpack = () => {
    const [pokemonData, setPokemonData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [isModalOpen, setIsModalOpen] = useState(false); // Modal visibility
    const [statusMessage, setStatusMessage] = useState(null); // Modal status message

    const fetchPokemonData = async () => {
        try {
            setLoading(true);
            const response = await fetch(
                `${import.meta.env.VITE_API_URL}/my-pokemon`
            );
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            setPokemonData(data.my_pokemons);
        } catch (error) {
            setError(error.message);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchPokemonData();
    }, []);

    const handleRefresh = () => {
        fetchPokemonData();
    };

    const handleModalClose = () => {
        setIsModalOpen(false);
        setStatusMessage(null); // Clear the message when closing the modal
    };

    const updatePokemonEntry = (updatedPokemon) => {
        setPokemonData((prevData) =>
            prevData.map((pokemon) =>
                pokemon.id === updatedPokemon.id ? updatedPokemon : pokemon
            )
        );
    };

    const handleRename = async (owned_id, newName) => {
        try {
            const response = await fetch(
                `${import.meta.env.VITE_API_URL}/my-pokemon/${owned_id}`,
                {
                    method: "PATCH",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({ nickname: newName }),
                }
            );
            const data = await response.json();
            if (data.success) {
                setStatusMessage({
                    isSuccess: true,
                    message: `Successfully renamed to ${newName}.`,
                });
                handleRefresh(); // Refresh the list in MyBackpack
            } else {
                setStatusMessage({
                    isSuccess: false,
                    message: data.message || "Failed to rename Pokémon.",
                });
            }
        } catch (error) {
            console.error("Error renaming Pokémon:", error);
            setStatusMessage({
                isSuccess: false,
                message: "An error occurred while renaming the Pokémon.",
            });
        } finally {
            setIsModalOpen(true);
        }
    };

    const handleRelease = async (owned_id) => {
        try {
            const response = await fetch(
                `${import.meta.env.VITE_API_URL}/my-pokemon/${owned_id}`,
                {
                    method: "DELETE",
                }
            );
            const data = await response.json();
            if (data.success) {
                setStatusMessage({
                    isSuccess: true,
                    message: `Successfully released Pokémon.`,
                });
                handleRefresh(); // Refresh the list in MyBackpack
            } else {
                setStatusMessage({
                    isSuccess: false,
                    message: "Failed to release Pokémon.",
                });
            }
        } catch (error) {
            console.error("Error releasing Pokémon:", error);
            setStatusMessage({
                isSuccess: false,
                message: "An error occurred while releasing the Pokémon.",
            });
        } finally {
            setIsModalOpen(true);
        }
    };

    if (loading) {
        return (
            <div className="flex justify-center items-center min-h-screen">
                <Spinner /> {/* Loading spinner */}
            </div>
        );
    }

    if (error) {
        return (
            <div className="flex justify-center items-center min-h-screen">
                <p className="text-red-500 text-lg">Error: {error}</p>
            </div>
        );
    }

    return (
        <div className="container mx-auto py-8">
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                {pokemonData.map((pokemon) => (
                    <Card
                        key={pokemon.id}
                        name={pokemon.nickname}
                        sprite={pokemon.pokemon.sprite}
                        owned_id={pokemon.id}
                        pokemon={pokemon}
                        isMyBackpack={true}
                        onRelease={() => handleRelease(pokemon.id)}
                        onRename={(newName) =>
                            handleRename(pokemon.id, newName)
                        }
                    />
                ))}
            </div>

            {/* Render Modal */}
            <Modal
                isOpen={isModalOpen}
                onClose={handleModalClose}
                statusMessage={statusMessage}
            />
        </div>
    );
};

export default MyBackpack;
