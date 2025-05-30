import type { User } from '$lib/types';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch }): Promise<{ users: User[] }> => {
  const response = await fetch(`/api/users`);

  if (!response.ok) {
    return { users: [] };
  }

  const users = await response.json();

  return {
    users
  };
};
