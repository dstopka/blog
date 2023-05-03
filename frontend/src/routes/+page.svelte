<script lang="ts">
	import PostsList from '$lib/Post/List.svelte';

	/** @type {import('./$types').PageData} */
	export let data: any;

	const { posts }: { posts: any[] } = data;
</script>

<PostsList {posts} />
