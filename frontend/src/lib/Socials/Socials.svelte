<script lang="ts">
    export { colorClass as color }

    let colorClass: string;
</script>

<div class="row">
    <a class="link {colorClass}" href="https://github.com/dstopka">
        <i class="fa-brands fa-github" />
    </a>
</div>

<style>
    .row {
        display: flex;
        flex-direction: row;
    }

    .dark {
        color: black;
    }

    .dimmed {
        color: rgb(74, 74, 74);
    }
    
    .light {
        color: rgb(252, 252, 252);
    }

    a {
        text-decoration: none;
    }
    
    @media (hover: hover) and (pointer: fine) {
		a:hover {
			color: rgb(95, 187, 224);
		}
	}
</style>