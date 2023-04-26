<script lang="ts">
	export let tags: any[];
</script>

<div class="tags">
	{#each tags as tag}
		<span class="tag">{tag}</span>
	{/each}
</div>

<style>
	.tags {
		display: flex;
		flex-wrap: wrap;
		row-gap: 0.5rem;
        column-gap: 0.5rem;
	}

	.tag {
		padding: 0.25rem 0.6rem;
        background-color: rgb(226, 226, 226);
		border-radius: 0.6rem;
		color: rgb(74, 74, 74);
		font-size: 0.8rem;
		vertical-align: middle;
		font-weight: 400;
        font-size: 0.9rem;
	}
</style>
