<script lang="ts">
	import Content from './PostContent.svelte';
	// @ts-ignore
	import Giscus from '@giscus/svelte';
	import Tags from './Tags.svelte';
	import DateInfo from './PostMeta.svelte';
	import HeaderLine from '$lib/components/HeaderLine.svelte';

	export let post: any;
</script>

<section class="post">
	<header class="wrap-lg">
		<div class="title">
			<h1>{post.title}</h1>
		</div>
		<div class="meta">
			<DateInfo publishedTime={post.publishedTime} updatedTime={post.updatedTime} />
		</div>
		<div class="cover">
			<img class="entry-cover" src={post.coverImageURL} alt="" />
		</div>
	</header>
	<section class="content">
		<div class="wrap-md">
			<Content content={post.content} />
		</div>
	</section>
	<footer>
		<div class="wrap-lg">
			<section class="tags">
				<div class="tags-header">
					<HeaderLine header="Tags" />
				</div>
				<Tags tags={post.tags} />
			</section>
			<section class="post-comments">
				<Giscus
					repo="dstopka/blog-discussions"
					repoId="R_kgDOJbkkow"
					category="Announcements"
					categoryId="DIC_kwDOJbkko84CWD9J"
					mapping="og:title"
					strict="0"
					reactionsEnabled="1"
					emitMetadata="0"
					theme="light"
					lang="en"
				/>
			</section>
		</div>
	</footer>
</section>

<style>
	.post {
		padding: 0 1.5rem;
	}

	h1 {
		color: var(--color-text-dark-headings);
		font-size: 3.5rem;
		font-weight: 700;
	}

	.wrap-lg {
		max-width: 55rem;
		margin: auto;
	}

	.wrap-md {
		margin: auto;
		max-width: 45rem;
	}

	.entry-cover {
		object-fit: cover;
		width: 100%;
		margin: 0.5rem 0;
		border-radius: 1rem;
	}

	.post {
		flex-direction: column;
		display: flex;
		margin-top: 1rem;
	}

	.meta {
		display: flex;
		flex-direction: row;
		font-size: 1rem;
		font-weight: 300;
		margin: 0.5rem 0;
	}

	.tags {
		margin: 1rem 0 0 0;
	}

	.tags-header {
		font-size: 1.75rem;
		margin: 1rem 0;
	}

	.content {
		width: 100%;
		text-align: justify;
		font-size: 1.125rem;
		line-height: 1.75rem;
		-webkit-font-smoothing: antialiased;
	}

	.post-comments {
		margin: 2rem 0;
	}

	@media (max-width: 768px) {
		.content {
			font-size: 1rem;
			line-height: 1.5rem;
		}

		.meta {
			font-size: 0.875rem;
		}

		h1 {
			font-size: 2.5rem;
		}
	}
</style>
