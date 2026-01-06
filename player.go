components {
  id: "player"
  component: "/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"basket\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/game.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.769231
    y: 0.918033
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"star\"\n"
  "mask: \"hazard\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 3.0\n"
  "      y: -4.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"Player Basket\"\n"
  "  }\n"
  "  data: 81.52117\n"
  "  data: 95.9189\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
