import { useRouter } from "next/router";
import React from "react";
import useSWR from "swr";
const fetcher = (url, args) => fetch(url, args).then((res) => res.json());

const Confirm = () => {
  const router = useRouter();
  const { data, error } = useSWR(
    [
      process.env.NEXT_PUBLIC_BACKEND_URL +
        `/api/generate?playlist=${router.query.playlist}`,
      {
        credentials: "include",
        mode: "cors",
      },
    ],
    fetcher
  );
  if (error) return "An error has occurred.";
  if (!data) return "Loading...";
  return (
    <div>
      <h1 className="font-medium text-6xl">Let's Confirm</h1>
    </div>
  );
};

export default Confirm;
