module "network" {
  source = "./modules/network"
  common_tags = {
   "Name" = "nbastats"
  }
}
