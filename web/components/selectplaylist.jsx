import React from "react";

const SelectPlaylist = ({ src, title, author }) => {
  return (
    <div className="flex flex-col bg-gray-100 gap-y-4 p-8">
      <img src={src} style={{ width: "300px", height: "300px" }}></img>
      <h2 className="text-3xl pt-4 self-start">{title}</h2>
      <h3 className="text-2xl text-gray-600 self-start">By {author}</h3>
    </div>
  );
};

export default SelectPlaylist;
