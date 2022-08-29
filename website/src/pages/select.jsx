import SelectPlaylist from "../components/selectplaylist";
import React from "react";

const Select = ({ playlists }) => {
  return (
    <div className="flex flex-col gap-y-8">
      <h1 className="font-medium text-6xl">Select your playlist</h1>
      <div className="flex flex-wrap gap-8">
        {playlists.map((playlist) => (
          <SelectPlaylist></SelectPlaylist>
        ))}
      </div>
    </div>
  );
};

export default Select;
