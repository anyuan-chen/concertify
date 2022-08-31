import { useRouter } from "next/router";
import React, { useEffect, useState } from "react";
import EditPlaylistItem from "../components/editPlaylistItem";
import useSWR from "swr";

const Confirm = () => {
  const [playlist, setPlaylist] = useState(undefined);
  const [editing, setEditing] = useState(false);
  const [editingId, setEditingId] = useState("");

  useEffect(() => {
    const generatePlaylist = async () => {
      const params = new URLSearchParams(window.location.search);
      const res = await fetch(
        process.env.NEXT_PUBLIC_BACKEND_URL +
          `/api/generate?playlist=${params.get("playlist")}`,
        {
          credentials: "include",
          mode: "cors",
        }
      ).then((res) => res.json());
      setPlaylist(res);
    };
    generatePlaylist();
    setEditing(true);
  }, []);

  if (playlist === undefined) return "Loading...";
  return (
    <div className="flex flex-col gap-y-16">
      {!editing && <h1 className="font-medium text-6xl">Let's Confirm</h1>}
      {editing && playlist && (
        <EditPlaylistItem
          playlistItem={playlist[0]}
          setEditing={setEditing}
        ></EditPlaylistItem>
      )}
    </div>
  );
};

export default Confirm;
