package dir

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFile_IsMusic_givenMp3(t *testing.T) {
	f := File{"foo.mp3"}

	require.True(t, f.IsMusic())
}

func TestFile_IsMusic_givenMp3Uppercased(t *testing.T) {
	f := File{"foo.MP3"}

	require.True(t, f.IsMusic())
}

func TestFile_IsMusic_givenNoExtension(t *testing.T) {
	f := File{"foo"}
	require.False(t, f.IsMusic())
}

func TestFile_IsPlaylist_givenM3u(t *testing.T) {
	f := File{"foo.m3u"}

	require.True(t, f.IsPlaylist())
}

func TestFile_IsPlaylist_givenM3uUppercased(t *testing.T) {
	f := File{"foo.M3U"}

	require.True(t, f.IsPlaylist())
}

func TestFile_IsPlaylist_givenNonMusicExtension(t *testing.T) {
	f := File{"foo.txt"}

	require.False(t, f.IsPlaylist())
}
