import React from "react";

const SelectPlaylist = ({ src, title, author }) => {
  return (
    <div className="flex flex-col bg-gray-100 gap-y-8 py-4">
      <img src={src} style={{ width: "300px", height: "300px" }}></img>
      <h2 className="text-3xl">{title}</h2>
      <h3 className="text-xl text-gray-600">By {author}</h3>
    </div>
  );
};

export default SelectPlaylist;
