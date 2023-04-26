<script lang="ts">
	import SubHeader from './SubHeader.svelte';
	import Tags from './Tags.svelte';

	export let post: any;
</script>

<article class="entry">
	<header class="entry-header">
		<h2 class="entry-title">
			<a href="/posts/{post.slug}">
				{post.title}
			</a>
		</h2>
		<div class="entry-meta">
			<SubHeader publishedTime={post.publishedTime} updatedTime={post.updatedTime} />
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
		padding-top: 1rem;
		margin-top: 0.5rem;
	}

	.entry:first-child {
		border-top: 1px solid transparent;
	}

	.entry-header {
		padding: 0;
		margin: 0;
	}

	.entry-header h2 a {
		text-decoration: none;
		color: inherit;
	}

	.entry-header h2 a:hover {
		color: var(--blue);
	}

	.entry-title {
		margin-top: 1rem;
		margin-bottom: 0.25rem;
		font-size: 2.5rem;
		margin: 0;
	}

	.entry-meta {
		margin: 0.25rem 0;
		display: flex;
		flex-direction: row;
		gap: 0.5rem;
		flex-wrap: wrap;
	}

	.entry-body {
		display: flex;
		flex-direction: row;
		margin: auto;
		margin: 0.5rem 0;
	}

	.entry-body .entry-summary {
		color: rgb(41, 41, 41);
		margin: 0.5rem 0;
	}

	.entry-body .entry-cover {
		margin-left: auto;
		padding-left: 1rem;
		margin-top: 0.5rem;
	}

	.entry-body img.entry-cover {
		object-fit: cover;
		width: 128px;
		height: 128px;
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

	.entry-footer a:hover {
		color: rgb(95, 187, 224);
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

		.entry-meta {
			flex-direction: column;
		}
	}
</style>
