<script lang="ts">
	import Fa from 'svelte-fa/src/fa.svelte';
	import { faArrowLeft, faArrowRight } from '@fortawesome/free-solid-svg-icons';

	import SubHeader from './DateInfo.svelte';
	import Tags from './Tags.svelte';

	export let post: any;
</script>

<article class="entry">
	<header class="entry-header">
		<SubHeader publishedTime={post.publishedTime} updatedTime={post.updatedTime} />
		<h2 class="entry-title">
			<a href="/posts/{post.slug}">
				{post.title}
			</a>
		</h2>
		<div class="entry-tags">
			{#if post.tags}
				<Tags tags={post.tags} />
			{/if}
		</div>
	</header>
	<div class="entry-body">
		<p class="entry-summary">{post.description}</p>
		<img class="entry-cover" src={post.coverImageURL} alt="" />
	</div>
	<footer class="entry-footer">
		<a href="/posts/{post.slug}">
			Read more
			<Fa icon={faArrowRight} />
		</a>
	</footer>
</article>

<style>
	h2 {
		color: var(--color-headers);
	}

	.entry {
		border-top: 1px solid var(--color-bg-gray);
		padding: 1rem 0;
	}

	.entry:first-child {
		border-top: 1px solid transparent;
	}

	.entry-header {
		margin: 0.5rem 0;
		padding-top: 1rem;
	}

	.entry-header a {
		text-decoration: none;
		color: inherit;
	}

	.entry-title {
		font-size: 2rem;
		line-height: normal;
		color: var(--color-text-dark-headings);
	}

	.entry-tags {
		display: flex;
		flex-direction: row;
		gap: 0.5rem;
		flex-wrap: wrap;
		line-height: normal;
	}

	.entry-tags {
		margin-top: 0.5rem;
	}

	.entry-body {
		display: flex;
		flex-direction: row;
		margin: auto;
		margin: 0.5rem 0;
		align-items: center;
	}

	.entry-summary {
		margin: 0.5rem 0;
		flex-grow: 1;
		font-size: 0.9rem;
	}

	.entry-cover {
		padding-left: 1rem;
		margin-bottom: auto;
		margin-top: 0.5rem;
	}

	img.entry-cover {
		object-fit: contain;
		width: 30%;
	}

	.entry-footer {
		margin: 1.5rem 0;
	}

	.entry-footer a {
		text-transform: uppercase;
		text-decoration: none;
		color: var(--color-text-dark-decorator);
		font-size: 0.8rem;
		font-weight: 400;
		line-height: normal;
	}

	@media (max-width: 768px) {
		.entry-body {
			display: flex;
			flex-direction: column-reverse;
			margin: auto;
		}

		.entry-body img.entry-cover {
			object-fit: cover;
			width: 100%;
			aspect-ratio: initial;
		}

		.entry-body .entry-cover {
			margin: 1rem 0 0.5rem 0;
			padding: 0;
		}
	}

	@media (hover: hover) and (pointer: fine) {
		.entry-footer a:hover {
			color: var(--color-text-dark-headings);
		}
	}
</style>
