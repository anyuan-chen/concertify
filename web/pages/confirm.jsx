import { useRouter } from "next/router";
import React from "react";
import useSWR from "swr";
const fetcher = (url, args) => fetch(url, args).then((res) => res.json());

const Confirm = () => {
  const router = useRouter();
  const { data, error } = useSWR(
    [
      process.env.NEXT_PUBLIC_BACKEND_URL +
        `/generate?playlist=${router.query.playlist}`,
      {
        credentials: "include",
        mode: "cors",
      },
    ],
    fetcher
  );
  return (
    <div>
      <h1>Let's Confirm</h1>
    </div>
  );
};

export default Confirm;
