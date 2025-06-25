import { useQuery } from "@tanstack/react-query";
import { API_URL } from "@/App"
import { fetchWithAuth } from "@/lib/utils";

interface UseGetFeedProps {
  since: string | null;
  until?: string | null;
  limit?: number | null;
  offset?: number | null;
  sort?: string | null;
  tags?: string | null;
  search?: string | null;
};

export const useGetFeed = ({
  since,
  until,
  limit,
  offset,
  sort,
  tags,
  search,
}: UseGetFeedProps) => {
  const query = useQuery({
    queryKey: [
      "tasks",
      since,
      until,
      limit,
      offset,
      sort,
      tags,
      search,
    ],
    queryFn: async () => {
      const params = new URLSearchParams();
      if (since) params.append('since', since);
      if (until) params.append('until', until);
      if (limit != null) params.append('limit', limit.toString());
      if (offset != null) params.append('offset', offset.toString());
      if (sort) params.append('sort', sort);
      if (tags) params.append('tags', tags);
      if (search) params.append('search', search);

      const response = await fetchWithAuth(`${API_URL}/feed?${params.toString()}`, {
        method: "GET",
      })

      if (!response.ok) {
        throw new Error("Failed to fetch feed");
      }

      const { data } = await response.json();

      return data;
    },
  });

  return query;
}