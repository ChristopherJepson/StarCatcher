components {
  id: "star"
  component: "/main/Scripts/star.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"star\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.2
    y: 0.221115
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"star\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "      z: 0.0029575448\n"
  "      w: 0.99999565\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"Star\"\n"
  "  }\n"
  "  data: 20.710339\n"
  "  data: 16.485594\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
