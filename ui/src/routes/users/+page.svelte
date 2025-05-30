<script lang="ts">
	import type { PageData } from './$types';

	const { data }: { data: PageData } = $props();
</script>

<div>
	<h1>Users:</h1>
	<ul>
		{#each data.users as user (user.id)}
			<li><a href={`/users/${user.id}`}>{user.name}</a></li>
		{/each}
	</ul>
</div>
