<script lang="ts">
	import SubHeader from './MetaDate.svelte';
	import Tags from './Tags.svelte';

	export let post: any;
</script>

<article class="entry">
	<header class="entry-header">
		<div class="entry-date">
			<SubHeader publishedTime={post.publishedTime} updatedTime={post.updatedTime} />
		</div>
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
		<a href="/posts/{post.slug}">Read more &rarr;</a>
	</footer>
</article>

<style>
	.entry {
		border-top: 1px solid rgb(243, 243, 243);
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
	}

	.entry-date,
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

	.entry-body .entry-summary {
		color: rgb(41, 41, 41);
		margin: 0.5rem 0;
		flex-grow: 1;
	}

	.entry-body .entry-cover {
		padding-left: 1rem;
		margin-top: 0.5rem;
	}

	.entry-body img.entry-cover {
		object-fit: contain;
		aspect-ratio: 1;
		max-width: 30%;
	}

	.entry-footer {
		margin: 0.5rem 0;
	}

	.entry-footer a {
		text-transform: uppercase;
		text-decoration: none;
		color: rgb(41, 41, 41);
		font-size: 0.8rem;
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
			height: 100%;
		}

		.entry-body .entry-cover {
			margin: 1rem 0 0.5rem 0;
			padding: 0;
		}

		.entry-date {
			flex-direction: column;
		}
	}

	@media (hover: hover) and (pointer: fine) {
		a:hover {
			color: rgb(95, 187, 224);
		}
	}
</style>
