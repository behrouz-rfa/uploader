# golang-even-driven-file-and-transcoder

The service is responsible for transcoding a source video to one or more result videos of various formats/resolutions/bitrate/etc.

Initial requirements
asynchronous transcoding (callback the results)

various vertical resolutions (2160p, 1080p, 720p, 480p, 360p) - no upscale

horizontaly scaling micro-services

s3-compatible storage integration (for source and results)

generate m3u8 playlist

Components
Error loading the extension!
Transcoding Flow
Error loading the extension!
Technical details
ffmpeg is used as a transcoding tool that allows us to easily integrate any transformations supported by it.

### TODO LIST

- [x] Upload File throw to the files module
- [x] Mino Storage
- [ ] ffmpege
- [ ] Even Driven
- [ ] CI/CD
- [ ] k8

****
