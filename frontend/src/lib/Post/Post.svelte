<script lang="ts">
	import Content from './Content.svelte';
	import Giscus from '@giscus/svelte'
	import Tags from './Tags.svelte';
	import DateInfo from './DateInfo.svelte';

	export let post: any;
</script>

<section class="post">
	<header>
		<div class="entry-meta">
			<DateInfo publishedTime={post.publishedTime} updatedTime={post.updatedTime} />
		</div>
		<h2 id="entry-title">{post.title}</h2>
	</header>
	<section class="post-content">
		<img class="entry-cover" src={post.coverImageURL} alt="" />
		<Content content={post.content} />
	</section>
	<section class="tags">
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
</section>

<style>

	header {
		margin: 1rem 0;
		color: var(--color-text-dark-headings);
	}

	.entry-cover {
		max-height: 400px;
		object-fit: cover;
		width: 100%;
		margin: 0.5rem 0;
	}

	.post {
		flex-direction: column;
		display: flex;
		margin-top: 1rem;
	}

	.entry-meta {
		display: flex;
		flex-direction: row;
		font-size: 0.9rem;
		font-weight: 400;
	}

	#entry-title {
		font-size: 2.5rem;
		line-height: normal;
	}

	.tags {
		padding: 2rem 0;
		border-bottom: 1px solid var(--color-bg-gray);
	}

	.post-content {
		margin: 0 0;
	}

	.post-comments {
		margin: 2rem 0;
	}
</style>
