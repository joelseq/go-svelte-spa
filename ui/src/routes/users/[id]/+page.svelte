<script lang="ts">
	import type { PageProps } from './$types';

	const { data }: PageProps = $props();
</script>

<div>
	<h1>User: {data.user.name}, ID: {data.user.id}</h1>
</div>
