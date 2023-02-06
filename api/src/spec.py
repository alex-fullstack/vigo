from pydantic import BaseModel

from models import SpecificationElement


class SpecificationEBMLHeader(BaseModel):
    specification_element:  SpecificationElement 
    ebml_version:  SpecificationElement 
    ebml_read_version:  SpecificationElement 
    ebml_max_id_length:  SpecificationElement 
    ebml_max_size_length:  SpecificationElement 
    doc_type:  SpecificationElement 
    doc_type_version:  SpecificationElement 
    doc_type_read_version:  SpecificationElement 


class SpecificationSeek(BaseModel):
    specification_element:  SpecificationElement 
    seek_id:  SpecificationElement 
    seek_position:  SpecificationElement 


class SpecificationSeekHead(BaseModel):
    specification_element:  SpecificationElement 
    seek:  list[SpecificationSeek] 


class SpecificationChapterTranslate(BaseModel):
    specification_element:  SpecificationElement 
    chapter_translate_edition_uid:  list[SpecificationElement] 
    chapter_translate_codec:  SpecificationElement 
    chapter_translate_id:  SpecificationElement 


class SpecificationInfo(BaseModel):
    specification_element:  SpecificationElement 
    segment_uid:  SpecificationElement 
    segment_filename:  SpecificationElement 
    prev_uid:  SpecificationElement 
    prev_filename:  SpecificationElement 
    next_uid:  SpecificationElement 
    next_filename:  SpecificationElement 
    segment_family:  list[SpecificationElement] 
    chapter_translate:  list[SpecificationChapterTranslate] 
    timecode_scale:  SpecificationElement 
    duration:  SpecificationElement 
    date_utc:  SpecificationElement 
    title:  SpecificationElement 
    muxing_app:  SpecificationElement 
    writing_app:  SpecificationElement 


class SpecificationSilentTracks(BaseModel):
    specification_element:  SpecificationElement 
    silent_track_number:  list[SpecificationElement] 


class SpecificationBlockMore(BaseModel):
    specification_element:  SpecificationElement 
    block_add_id:  SpecificationElement 
    block_additional:  SpecificationElement 


class SpecificationBlockAdditions(BaseModel):
    specification_element:  SpecificationElement 
    block_more:  list[SpecificationBlockMore] 


class SpecificationTimeSlice(BaseModel):
    specification_element:  SpecificationElement 
    lace_number:  SpecificationElement 
    frame_number:  SpecificationElement 
    block_addition_id:  SpecificationElement 
    delay:  SpecificationElement 
    slice_duration:  SpecificationElement 


class SpecificationSlices(BaseModel):
    specification_element:  SpecificationElement 
    time_slice:  list[SpecificationTimeSlice] 


class SpecificationReferenceFrame(BaseModel):
    specification_element:  SpecificationElement 
    reference_offset:  SpecificationElement 
    reference_time_code:  SpecificationElement 


class SpecificationBlockGroup(BaseModel):
    specification_element:  SpecificationElement 
    block:  SpecificationElement 
    block_virtual:  SpecificationElement 
    block_additions:  SpecificationBlockAdditions 
    block_duration:  SpecificationElement 
    reference_priority:  SpecificationElement 
    reference_block:  list[SpecificationElement] 
    reference_virtual:  SpecificationElement 
    codec_state:  SpecificationElement 
    discard_padding:  SpecificationElement 
    slices:  SpecificationSlices 
    reference_frame:  SpecificationReferenceFrame 


class SpecificationCluster(BaseModel):
    specification_element:  SpecificationElement 
    timecode:  SpecificationElement 
    silent_tracks:  SpecificationSilentTracks 
    position:  SpecificationElement 
    prev_size:  SpecificationElement 
    simple_block:  list[SpecificationElement] 
    block_group:  list[SpecificationBlockGroup] 
    encrypted_block:  list[SpecificationElement] 


class SpecificationTrackTranslate(BaseModel):
    specification_element:  SpecificationElement 
    track_translate_edition_uid:  list[SpecificationElement] 
    track_translate_codec:  SpecificationElement 
    track_translate_track_id:  SpecificationElement 


class SpecificationVideo(BaseModel):
    specification_element:  SpecificationElement 
    flag_interlaced:  SpecificationElement 
    stereo_mode:  SpecificationElement 
    alpha_mode:  SpecificationElement 
    old_stereo_mode:  SpecificationElement 
    pixel_width:  SpecificationElement 
    pixel_height:  SpecificationElement 
    pixel_crop_bottom:  SpecificationElement 
    pixel_crop_top:  SpecificationElement 
    pixel_crop_left:  SpecificationElement 
    pixel_crop_right:  SpecificationElement 
    display_width:  SpecificationElement 
    display_height:  SpecificationElement 
    display_unit:  SpecificationElement 
    aspect_ratio_type:  SpecificationElement 
    colour_space:  SpecificationElement 
    gamma_value:  SpecificationElement 
    frame_rate:  SpecificationElement 


class SpecificationAudio(BaseModel):
    specification_element:  SpecificationElement 
    sampling_frequency:  SpecificationElement 
    output_sampling_frequency:  SpecificationElement 
    channels:  SpecificationElement 
    channel_positions:  SpecificationElement 
    bit_depth:  SpecificationElement 


class SpecificationTrackPlane(BaseModel):
    specification_element:  SpecificationElement 
    track_plane_uid:  SpecificationElement 
    track_plane_type:  SpecificationElement 


class SpecificationTrackCombinePlanes(BaseModel):
    specification_element:  SpecificationElement 
    track_plane:  list[SpecificationTrackPlane] 


class SpecificationTrackJoinBlocks(BaseModel):
    specification_element:  SpecificationElement 
    track_join_uid:  list[SpecificationElement] 


class SpecificationTrackOperation(BaseModel):
    specification_element:  SpecificationElement 
    track_combine_planes:  SpecificationTrackCombinePlanes 
    track_join_blocks:  SpecificationTrackJoinBlocks 


class SpecificationContentCompression(BaseModel):
    specification_element:  SpecificationElement 
    content_comp_algo:  SpecificationElement 
    content_comp_settings:  SpecificationElement 


class SpecificationContentEncryption(BaseModel):
    specification_element:  SpecificationElement 
    content_enc_algo:  SpecificationElement 
    content_enc_key_id:  SpecificationElement 
    content_signature:  SpecificationElement 
    content_sig_key_id:  SpecificationElement 
    content_sig_algo:  SpecificationElement 
    content_sig_hash_algo:  SpecificationElement 


class SpecificationContentEncoding(BaseModel):
    specification_element:  SpecificationElement 
    content_encoding_order:  SpecificationElement 
    content_encoding_scope:  SpecificationElement 
    content_encoding_type:  SpecificationElement 
    content_compression:  SpecificationContentCompression 
    content_encryption:  SpecificationContentEncryption 


class SpecificationContentEncodings(BaseModel):
    specification_element:  SpecificationElement 
    content_encoding:  list[SpecificationContentEncoding] 


class SpecificationTrackEntry(BaseModel):
    specification_element:  SpecificationElement 
    track_number:  SpecificationElement 
    track_uid:  SpecificationElement 
    track_type:  SpecificationElement 
    flag_enabled:  SpecificationElement 
    flag_default:  SpecificationElement 
    flag_forced:  SpecificationElement 
    flag_lacing:  SpecificationElement 
    min_cache:  SpecificationElement 
    max_cache:  SpecificationElement 
    default_duration:  SpecificationElement 
    default_decoded_field_duration:  SpecificationElement 
    track_timecode_scale:  SpecificationElement 
    track_offset:  SpecificationElement 
    max_block_addition_id:  SpecificationElement 
    name:  SpecificationElement 
    language:  SpecificationElement 
    codec_id:  SpecificationElement 
    codec_private:  SpecificationElement 
    codec_name:  SpecificationElement 
    attachment_link:  SpecificationElement 
    codec_settings:  SpecificationElement 
    codec_info_url:  list[SpecificationElement] 
    codec_download_url:  list[SpecificationElement] 
    codec_decode_all:  SpecificationElement 
    track_overlay:  list[SpecificationElement] 
    codec_delay:  SpecificationElement 
    seek_pre_roll:  SpecificationElement 
    track_translate:  list[SpecificationTrackTranslate] 
    video:  SpecificationVideo 
    audio:  SpecificationAudio 
    track_operation:  SpecificationTrackOperation 
    trick_track_uid:  SpecificationElement 
    trick_track_segment_uid:  SpecificationElement 
    trick_track_flag:  SpecificationElement 
    trick_master_track_uid:  SpecificationElement 
    trick_master_track_segment_uid:  SpecificationElement 
    content_encodings:  SpecificationContentEncodings 


class SpecificationTracks(BaseModel):
    specification_element:  SpecificationElement 
    track_entry:  list[SpecificationTrackEntry] 


class SpecificationCueReference(BaseModel):
    specification_element:  SpecificationElement 
    cue_ref_time:  SpecificationElement 
    cue_ref_cluster:  SpecificationElement 
    cue_ref_number:  SpecificationElement 
    cue_ref_codec_state:  SpecificationElement 


class SpecificationCueTrackPositions(BaseModel):
    specification_element:  SpecificationElement 
    cue_track:  SpecificationElement 
    cue_cluster_position:  SpecificationElement 
    cue_relative_position:  SpecificationElement 
    cue_duration:  SpecificationElement 
    cue_block_number:  SpecificationElement 
    cue_codec_state:  SpecificationElement 
    cue_reference:  list[SpecificationCueReference] 


class SpecificationCuePoint(BaseModel):
    specification_element:  SpecificationElement 
    cue_time:  SpecificationElement 
    cue_track_positions:  list[SpecificationCueTrackPositions] 


class SpecificationCues(BaseModel):
    specification_element:  SpecificationElement 
    cue_point:  list[SpecificationCuePoint] 


class SpecificationAttachedFile(BaseModel):
    specification_element:  SpecificationElement 
    file_description:  SpecificationElement 
    file_name:  SpecificationElement 
    file_mime_type:  SpecificationElement 
    file_data:  SpecificationElement 
    file_uid:  SpecificationElement 
    file_referral:  SpecificationElement 
    file_used_start_time:  SpecificationElement 
    file_used_end_time:  SpecificationElement 


class SpecificationAttachments(BaseModel):
    specification_element:  SpecificationElement 
    attached_file:  list[SpecificationAttachedFile] 


class SpecificationChapterTrack(BaseModel):
    specification_element:  SpecificationElement 
    chapter_track_number:  list[SpecificationElement] 


class SpecificationChapterDisplay(BaseModel):
    specification_element:  SpecificationElement 
    chap_string:  SpecificationElement 
    chap_language:  list[SpecificationElement] 
    chap_country:  list[SpecificationElement] 


class SpecificationChapProcessCommand(BaseModel):
    specification_element:  SpecificationElement 
    chap_process_time:  SpecificationElement 
    chap_process_data:  SpecificationElement 


class SpecificationChapProcess(BaseModel):
    specification_element:  SpecificationElement 
    chap_process_codec_id:  SpecificationElement 
    chap_process_private:  SpecificationElement 
    chap_process_command:  list[SpecificationChapProcessCommand] 


class SpecificationChapterAtom(BaseModel):
    specification_element:  SpecificationElement 
    chapter_uid:  SpecificationElement 
    chapter_string_uid:  SpecificationElement 
    chapter_time_start:  SpecificationElement 
    chapter_time_end:  SpecificationElement 
    chapter_flag_hidden:  SpecificationElement 
    chapter_flag_enabled:  SpecificationElement 
    chapter_segment_uid:  SpecificationElement 
    chapter_segment_edition_uid:  SpecificationElement 
    chapter_physical_equiv:  SpecificationElement 
    chapter_track:  SpecificationChapterTrack 
    chapter_display:  list[SpecificationChapterDisplay] 
    chap_process:  list[SpecificationChapProcess] 


class SpecificationEditionEntry(BaseModel):
    specification_element:  SpecificationElement 
    edition_uid:  SpecificationElement 
    edition_flag_hidden:  SpecificationElement 
    edition_flag_default:  SpecificationElement 
    edition_flag_ordered:  SpecificationElement 
    chapter_atom:  list[SpecificationChapterAtom] 


class SpecificationChapters(BaseModel):
    specification_element:  SpecificationElement 
    edition_entry:  list[SpecificationEditionEntry] 


class SpecificationTargets(BaseModel):
    specification_element:  SpecificationElement 
    target_type_value:  SpecificationElement 
    target_type:  SpecificationElement 
    tag_track_uid:  list[SpecificationElement] 
    tag_edition_uid:  list[SpecificationElement] 
    tag_chapter_uid:  list[SpecificationElement] 
    tag_attachment_uid:  list[SpecificationElement] 


class SpecificationSimpleTag(BaseModel):
    specification_element:  SpecificationElement 
    tag_name:  SpecificationElement 
    tag_language:  SpecificationElement 
    tag_default:  SpecificationElement 
    tag_string:  SpecificationElement 
    tag_binary:  SpecificationElement 


class SpecificationTag(BaseModel):
    specification_element:  SpecificationElement 
    targets:  SpecificationTargets 
    simple_tag:  list[SpecificationSimpleTag] 


class SpecificationTags(BaseModel):
    specification_element:  SpecificationElement 
    tag:  list[SpecificationTag] 


class SpecificationSegment(BaseModel):
    specification_element:  SpecificationElement 
    ebml_empty:  SpecificationElement 
    seek_head:  list[SpecificationSeekHead] 
    info:  list[SpecificationInfo] 
    cluster:  list[SpecificationCluster] 
    tracks:  list[SpecificationTracks] 
    cues:  SpecificationCues 
    attachments:  SpecificationAttachments 
    chapters:  SpecificationChapters 
    tags:  list[SpecificationTags] 


class Specification(BaseModel):
    ebml_header:  list[SpecificationEBMLHeader] 
    segment:  list[SpecificationSegment] 


