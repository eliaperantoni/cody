<template>
    <div class="footer">
        Made with <span class="love">♥</span> by 
        <span class="at">@eliaperantoni</span> and 
        <span class="at">@nicofretti</span>
    </div>
</template>

<script>
export default {
    name: "footer",
}
</script>

<style lang="scss" scoped>
.footer {
    user-select: none;
    font-family: "Fira Code";
    color: rgb(122, 122, 122);
    span.love {
        color: #ff295b;
        font-size:  1.5em;
        text-shadow: 0 0 3px #ff295b;
    }
    span.at {
        text-decoration: underline;
        cursor: pointer;
    }
}
</style>
