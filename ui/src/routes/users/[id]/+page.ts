import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, params }) => {
  const { id } = params;

  if (!id) {
    return { user: null };
  }

  const response = await fetch(`/api/users/${id}`);

  if (!response.ok) {
    return { user: null };
  }

  const user = await response.json();
  return { user };
};
