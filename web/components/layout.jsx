import React from "react";
import Link from "next/link";
const Layout = ({ children }) => {
  return (
    <div className="w-screen lg:px-32 lg:py-16">
      <div className="w-full flex justify-center flex-col gap-y-24">
        <Link href="/">
          <a className="font-medium text-3xl">Concertify</a>
        </Link>
        <div className="xlg:w-[1400px]">{children}</div>
      </div>
    </div>
  );
};

export default Layout;
