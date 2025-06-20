import { useMutation, useQueryClient } from "@tanstack/react-query";
import { toast } from "sonner";
import { useNavigate } from "react-router-dom";

import { API_URL } from "@/App"

interface RequestType {
  json: Record<string, any>; // Define JSON payload structure
}

interface ResponseType {
  token: string; // Adjust based on API response structure
}

export const useLogin = () => {
  //const queryClient = useQueryClient();

  const navigate = useNavigate();

  const mutation = useMutation<ResponseType, Error, RequestType>({
    mutationFn: async ({ json }) => {
      const response = await fetch(`${API_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(json),
      });

      if (!response.ok) {
        throw new Error("Failed to log in");
      }

      return await response.json();
    },
    onSuccess: (data) => {
      localStorage.setItem("token", data.token); // Store token for authentication
      toast.success("Logged in successfully!");
      //queryClient.invalidateQueries({ queryKey: ["current"] }); // Refresh session
      navigate("/feed");
    },
    onError: () => {
      toast.error("Login failed. Please check your credentials.");
    },
  });

  return mutation;
};
