import { SearchService } from "./search.service";
import { HTTP_INTERCEPTORS, HttpClient, HttpClientModule } from "@angular/common/http";

// Straight Jasmine testing without Angular's testing support
describe('SearchService', () => {
    let service: SearchService;
    // beforeEach(() => {
    //   service = new SearchService(new HttpClient());
    // });
  
    it('#getValue should return real value', () => {
       let profile = service.searchProfile("Miposhka","7.35");
       console.log(profile);
    });
  });