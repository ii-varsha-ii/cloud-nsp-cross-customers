Cross Customers:

# prelims 
Backend:
1. if organization does not exist:
   2. create organization
2. create new account / add existing account
   3. use AWS SDK for this
   4. Send out invitation
   5. Use refresh to check whether account is added in organization
3. select scenario and region
4. Create TGW in that region and check whether the accounts are able to connect to that TGW


Folder structure:

1. Organizations and Accounts
2. Terraform for creating an organization and setting permissions & check existing        
<br>
1. VPCs
2. Terraform for creating new VPC or selecting existing VPC    
<br>
1. TGW
2. Terraform for creating/selecting existing TGWs and creating TGW attachments & check existing   
<br>
1. Routing Table configs
2. Terraform for creating mappings between two resources.    
<br>
1. VMs
2. Terrform for creating VMs


Routers:  
**org_router.go**
1. Create Organization - (do it manually)
2. Get Organization based on ID
3. Create Accounts in Organization
4. Invite Account to Organization
5. Get all Accounts in Organization
6. Get Account based on AccountID

**vpc_router.go**
1. Create VPC in a region
2. Get VPC in a region
3. Create Subnets in a VPC   
   3.i Number of public subnets   
   3.ii Number of private subnets
4. Get all subnets for a VPC in a region

**common_router.go**
1. Connect two VPCs
Based on Region:
   4.1 create transit gateway   
   4.2 create two transit gateway attachments   
   4.3 update routes to transit gateway
   4.4 connect done? yes / no

**gateway_router.go**
1. Get Transit Gateway in region
   1.1 Create Transit Gateway in that region

**vm_router.go**
1. Create an EC2 given VPC given region


20/3 Monday:

Flow:
1. User comes to the page and ask his account to be attached to the organization
   a. get the list of orgs -> select organization
   b. get the list of accounts -> select his account
   c. create his account to the organization
      c.1 either send invitation / create an account
   d. verify by getting his account after creation
2. select vpc & subnets / create vpc & subnets


Generate GRPC code:
`protoc \
--go_out=./backend/go_pb_files/ \
--go_opt=module=github.com/ii-varsha-ii/nsp-cross-customer-multicloud/backend/protos \
--go-grpc_out=./backend/go_pb_files/ \
--go-grpc_opt=module=github.com/ii-varsha-ii/nsp-cross-customer-multicloud/backend/protos \
./backend/protos/organization.proto`