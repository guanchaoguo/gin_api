-- 关键词视图
CREATE VIEW `yyets2019`.`resoure_search`  AS
(SELECT
  rid origin_id,
  poster,
  cnname,
  enname,
  type_id,
  '' metainfo,
  favorites,
  '' author,
  update_time
FROM
  resource );