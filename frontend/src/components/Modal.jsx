import React from "react";
import { FaTimes } from "react-icons/fa"; // Import the close icon

const Modal = ({ isOpen, onClose, statusMessage, children }) => {
    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
            <div className="bg-white p-8 rounded-lg shadow-lg text-center relative">
                {/* Close Icon */}
                <button
                    className="absolute top-4 right-4 text-gray-500 hover:text-gray-700"
                    onClick={onClose}
                >
                    <FaTimes size={24} />
                </button>

                {/* Modal Content */}
                {children}
                {statusMessage && (
                    <div
                        className={`mt-4 ${
                            statusMessage.isSuccess
                                ? "text-green-500"
                                : "text-red-500"
                        }`}
                    >
                        {statusMessage.message}
                    </div>
                )}
            </div>
        </div>
    );
};

export default Modal;
