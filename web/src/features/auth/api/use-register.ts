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

export const useRegister = () => {
  //const queryClient = useQueryClient();

  const navigate = useNavigate();

  const mutation = useMutation<
    ResponseType,
    Error,
    RequestType
  >({
    mutationFn: async ({ json }) => {
      const response = await fetch(`${API_URL}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(json),
      });

      if (!response.ok) {
        throw new Error("Failed to register");
      }

      return await response.json();
    },
    onSuccess: (data) => {
      localStorage.setItem("token", data.token);
      toast.success("Registered");
      //queryClient.invalidateQueries({ queryKey: ["current"] });
      navigate("/feed");
    },
    onError: () => {
      toast.error("Failed to register");
    },
  });

  return mutation;
};
