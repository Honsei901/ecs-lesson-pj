import React from 'react';
import { healthCheck } from '../api/apiClient';

interface Props {
  setText: React.Dispatch<React.SetStateAction<string>>;
}

export const HealthCheckButton = ({ setText }: Props) => {
  const getText = async () => {
    const response = await healthCheck();
    setText(response as string);
  };
  return (
    <button
      className="bg-blue-500 p-5 rounded-md font-bold text-lg text-white"
      onClick={getText}
    >
      HealthCheckButton
    </button>
  );
};
