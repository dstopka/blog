<script lang="ts">
	import Fa from 'svelte-fa';
	import { faArrowRight } from '@fortawesome/free-solid-svg-icons';

	import Meta from './PostMeta.svelte';
	import Tags from './Tags.svelte';

	export let post: any;
</script>

<article class="entry">
	<header>
		<h1>
			<a href="/posts/{post.slug}">
				{post.title}
			</a>
		</h1>
		<div class="subheader">
			<Meta publishedTime={post.publishedTime} updatedTime={post.updatedTime} />
			{#if post.tags}
				<Tags tags={post.tags} />
			{/if}
		</div>
	</header>
	<section class="entry-body">
		<p class="entry-summary">{post.description}</p>
		<a class="entry-cover" href="/posts/{post.slug}">
			<img src={post.coverImageURL} alt="" />
		</a>
	</section>
	<footer>
		<a href="/posts/{post.slug}">
			Read more
			<div class="btn-arrow">
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

	.entry:first-child header {
		margin: 0;
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

	h1 {
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
		gap: 1rem;
	}

	.entry-summary {
		flex-grow: 1;
		font-size: 1rem;
		-webkit-font-smoothing: antialiased;
		line-height: 1.5rem;
		flex: 1 1 0px;
	}

	.entry-cover {
		margin-bottom: auto;
		flex: 1 1 0px;
		max-width: 40%;
	}

	.entry-cover img {
		object-fit: cover;
		width: 100%;
		border-radius: 0.5rem;
	}

	footer {
		margin-top: 1.5rem;
	}

	footer a {
		background-color: var(--color-read-more-button-bg);
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
		h1 {
			font-size: 1.5rem;
		}

		.entry-body {
			display: flex;
			flex-direction: column-reverse;
			gap: 0.5rem;
			margin-top: 1rem;
		}

		.entry-body .entry-cover {
			width: 100%;
		}

		.entry-cover {
			max-width: 100%;
			padding: 0;
		}

		.entry-cover img {
			aspect-ratio: initial;
		}
	}

	@media (hover: hover) and (pointer: fine) {
		footer a:hover {
			background-color: var(--color-read-more-button-bg-hover);
		}

		footer a:hover .btn-arrow {
			transform: translateX(3px);
		}
	}
</style>
