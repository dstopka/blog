<script lang="ts">
	import { formatDate } from '$lib/date';

	export let publishedTime: string;
	export let updatedTime: string;
</script>

<div class="sub-header">
	<span class="sub-header-date">
		Posted on {formatDate(publishedTime)}
	</span>
	{#if updatedTime}
		<span class="sub-header-date">•</span>
		<span class="sub-header-date">updated on {formatDate(updatedTime)}</span>
	{/if}
</div>

<style>
	.sub-header {
		display: flex;
		flex-direction: row;
		font-size: 1rem;
		font-weight: 400;
		color: rgb(74, 74, 74);
        gap: 0.25rem;
	}

	.sub-header-date {
		margin: 0.25rem 0;
		flex-shrink: 0;
		text-align: center;
	}
</style>
