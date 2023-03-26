import { useContext } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { UserContext } from "../context/userContext";

export default function CustomerRoute() {
  const [state] = useContext(UserContext);
  return state.isLogin === true ? <Outlet /> : <Navigate to="/" />;
}