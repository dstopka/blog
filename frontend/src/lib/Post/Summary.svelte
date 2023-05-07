<script lang="ts">
	import Fa from 'svelte-fa/src/fa.svelte';
	import { faArrowRight } from '@fortawesome/free-solid-svg-icons';

	import Meta from './Meta.svelte';
	import Tags from './Tags.svelte';

	export let post: any;
</script>

<article class="entry">
	<header>
		<h2>
			<a href="/posts/{post.slug}">
				{post.title}
			</a>
		</h2>
		<div class="subheader">
			<Meta publishedTime={post.publishedTime} updatedTime={post.updatedTime} />
			{#if post.tags}
				<Tags tags={post.tags} />
			{/if}
		</div>
	</header>
	<section class="entry-body">
		<p class="entry-summary">{post.description}</p>
		<img class="entry-cover" src={post.coverImageURL} alt="" />
	</section>
	<footer>
		<a href="/posts/{post.slug}">
			Read more
			<div class="btn-arrow" >
				<Fa icon={faArrowRight} />
			</div>
		</a>
	</footer>
</article>

<style>
	.entry {
		border-top: 1px solid var(--color-bg-gray);
		padding: 1rem 0;
	}

	.entry:first-child {
		border-top: 1px solid transparent;
		padding-top: 0;
	}

	header {
		margin: 0.5rem 0;
		padding-top: 1rem;
	}

	header a {
		text-decoration: none;
		color: inherit;
	}

	h2 {
		font-size: 2rem;
		line-height: normal;
		color: var(--color-text-dark-headings);
	}

	.subheader {
		display: flex;
		flex-direction: row;
		align-items: center;
		flex-wrap: wrap;
		gap: 0.5rem;
		margin: 0.5rem 0;
		font-weight: 300;
		font-size: 0.875rem;
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
		font-size: 1.125rem;
		-webkit-font-smoothing: antialiased;
		line-height: 1.75rem;
	}

	.entry-cover {
		margin-left: 1rem;
		margin-bottom: auto;
		margin-top: 0.5rem;
		object-fit: cover;
		width: 50%;
		border-radius: 1rem;
	}

	footer {
		margin-top: 1.5rem;
	}

	footer a {
		background-color: var(--color-logo-secondary-bg);
		text-transform: uppercase;
		text-decoration: none;
		color: var(--color-text-light-primary);
		font-size: 0.75rem;
		font-weight: 400;
		line-height: normal;
		padding: 0.5rem 1rem;
		border-radius: 1rem;
	}

	.btn-arrow {
		display: inline-block;
		transition: all 0.2s ease-in-out;
	}

	@media (max-width: 768px) {
		.subheader {
			flex-direction: column;
			align-items: start;
		}

		.entry-body {
			display: flex;
			flex-direction: column-reverse;
			margin: auto;
		}

		.entry-body img.entry-cover {
			width: 100%;
			aspect-ratio: initial;
		}

		.entry-body .entry-cover {
			margin: 1rem 0 0.5rem 0;
			padding: 0;
		}
	}

	@media (hover: hover) and (pointer: fine) {
		footer a:hover {
			background-color: var(--color-logo-secondary-bg-hover);
		}

		footer a:hover .btn-arrow {
			transform: translateX(3px);
		}
	}
</style>
