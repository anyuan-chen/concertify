import React from "react";
import YoutubeVideoPreview from "./youtubeVideoPreview";

const PlaylistItemPreview = ({ playlistItem, setEditing, setEditingId }) => {
  return (
    <div className="border border-black p-4  flex flex-col gap-y-8">
      <div className="flex gap-x-4 items-end">
        <h3 className="text-3xl font-medium">
          {playlistItem.name.replace(/&quot;/g, '\\"')}
        </h3>
        <h3 className="text-xl text-grey-600">
          {playlistItem.artists
            .reduce((prev, cur) => {
              return [...prev, cur.name];
            }, "")
            .join(", ")}
        </h3>
      </div>
      <YoutubeVideoPreview
        href={`https://www.youtube.com/watch?v=${playlistItem.youtube_search_response[0].id.videoId}`}
        src={
          playlistItem.youtube_search_response[0].snippet.thumbnails.default.url
        }
        title={playlistItem.youtube_search_response[0].snippet.title}
        artist={playlistItem.youtube_search_response[0].snippet.channelTitle}
      ></YoutubeVideoPreview>
      <button
        className="self-start text-xl underline"
        onClick={() => {
          setEditingId(playlistItem.id);
        }}
      >
        choose another performance
      </button>
    </div>
  );
};

export default PlaylistItemPreview;
