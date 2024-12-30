import React from 'react';
import { checkRestApi } from '../api/apiClient';

interface Props {
  setText: React.Dispatch<React.SetStateAction<string>>;
}

export const ApiButton = ({ setText }: Props) => {
  const getText = async () => {
    const response = await checkRestApi();
    setText(response as string);
  };
  return (
    <button
      className="bg-red-500 p-5 rounded-md font-bold text-lg text-white"
      onClick={getText}
    >
      ApiButton
    </button>
  );
};
