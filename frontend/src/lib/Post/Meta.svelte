<script lang="ts">
	import { formatDate } from '$lib/date';

	export let publishedTime: string;
	export let updatedTime: string;
</script>

<div class="meta">
	<span class="meta-date">
		{formatDate(publishedTime)}
	</span>
	{#if updatedTime}
		<span class="meta-date">•</span>
		<span class="meta-date">updated on {formatDate(updatedTime)}</span>
	{/if}
</div>

<style>
	.meta {
		display: flex;
		flex-direction: row;
		color: var(--color-text-dark-primary);
		gap: 0.25rem;
	}

	.meta-date {
		flex-shrink: 0;
		text-align: center;
	}
</style>
