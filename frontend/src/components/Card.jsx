import React, { useState } from "react";
import { useNavigate } from "react-router-dom"; // Import useNavigate hook
import Modal from "./Modal"; // Import the Modal component

const Card = ({
    name,
    sprite,
    owned_id,
    id,
    url,
    onRelease,
    onRename,
    isMyBackpack = false, // Flag to differentiate between contexts
}) => {
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [modalType, setModalType] = useState(null); // Track the modal type (rename or release)
    const [statusMessage, setStatusMessage] = useState(null); // Track the response status message
    const [newName, setNewName] = useState(""); // State for new name input
    const [isLoading, setIsLoading] = useState(false); // Loading state for submit actions

    const navigate = useNavigate(); // Initialize useNavigate hook

    const handleModalClose = () => {
        setIsModalOpen(false);
        setModalType(null);
        setStatusMessage(null); // Clear the message when closing the modal
        setNewName(""); // Clear the new name input
        setIsLoading(false); // Reset loading state
    };

    const handleRenameSubmit = async () => {
        if (!newName.trim()) return; // Prevent renaming to empty name
        setIsLoading(true);
        try {
            await onRename(newName); // Call the prop function
        } catch (error) {
            console.error("Error renaming Pokémon:", error);
            setStatusMessage({
                isSuccess: false,
                message: "An error occurred while renaming the Pokémon.",
            });
        } finally {
            setIsLoading(false);
            setIsModalOpen(true); // Show modal after rename operation
        }
    };

    const handleRelease = async () => {
        try {
            await onRelease(); // Call the prop function
        } catch (error) {
            console.error("Error releasing Pokémon:", error);
            setStatusMessage({
                isSuccess: false,
                message: "An error occurred while releasing the Pokémon.",
            });
        } finally {
            setIsModalOpen(true); // Show modal after release operation
        }
    };

    const handleViewDetails = () => {
        navigate(`/pokemon/${id}`); // Redirect to PokémonDetail page
    };

    return (
        <div className="bg-white shadow-lg rounded-lg p-4">
            <img src={sprite} alt={name} className="w-24 h-24 mx-auto" />
            <h3 className="text-xl text-center mt-4">{name}</h3>
            <div className="flex justify-between mt-4">
                {isMyBackpack ? (
                    <>
                        <button
                            className="bg-blue-500 text-white py-1 px-3 rounded"
                            onClick={() => {
                                setModalType("rename");
                                setIsModalOpen(true);
                            }}
                        >
                            Rename
                        </button>
                        <button
                            className="bg-red-500 text-white py-1 px-3 rounded"
                            onClick={() => {
                                setModalType("release");
                                setIsModalOpen(true);
                            }}
                        >
                            Release
                        </button>
                    </>
                ) : (
                    <button
                        className="bg-blue-500 text-white py-1 px-3 rounded"
                        onClick={handleViewDetails} // Handle redirect to PokémonDetail
                    >
                        View Details
                    </button>
                )}
            </div>

            {/* Render Modal */}
            {isModalOpen && (
                <Modal
                    isOpen={isModalOpen}
                    onClose={handleModalClose}
                    statusMessage={statusMessage}
                >
                    {modalType === "rename" && (
                        <>
                            <h2 className="text-2xl font-bold mb-4">
                                Rename Pokémon
                            </h2>
                            <input
                                type="text"
                                placeholder="Enter new nickname"
                                className="border border-gray-300 rounded p-2 mb-4 w-full"
                                value={newName}
                                onChange={(e) => setNewName(e.target.value)}
                                onKeyDown={(e) => {
                                    if (e.key === "Enter") {
                                        handleRenameSubmit();
                                    }
                                }}
                            />
                            <div className="flex justify-between mt-4">
                                <button
                                    className="bg-green-500 text-white py-2 px-4 rounded"
                                    onClick={handleRenameSubmit}
                                    disabled={isLoading}
                                >
                                    {isLoading ? "Renaming..." : "Submit"}
                                </button>
                                <button
                                    className="bg-gray-500 text-white py-2 px-4 rounded"
                                    onClick={handleModalClose}
                                >
                                    Cancel
                                </button>
                            </div>
                        </>
                    )}
                    {modalType === "release" && (
                        <>
                            <h2 className="text-2xl font-bold mb-4 text-red-500">
                                Are you sure you want to release {name}?
                            </h2>
                            <div className="flex justify-between mt-4">
                                <button
                                    className="bg-red-500 text-white py-2 px-4 rounded"
                                    onClick={handleRelease}
                                    disabled={isLoading}
                                >
                                    {isLoading ? "Releasing..." : "Release"}
                                </button>
                                <button
                                    className="bg-gray-500 text-white py-2 px-4 rounded"
                                    onClick={handleModalClose}
                                >
                                    Cancel
                                </button>
                            </div>
                        </>
                    )}
                </Modal>
            )}
        </div>
    );
};

export default Card;
