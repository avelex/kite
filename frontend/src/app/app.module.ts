import { APP_INITIALIZER, NgModule } from "@angular/core";
import { BrowserModule } from "@angular/platform-browser";

import { AppComponent } from "./app.component";
import { AppRoutingModule } from "./app-routing.module";
import { FooterComponent } from "./core/layout/footer.component";
import { HeaderComponent } from "./core/layout/header.component";
import { EMPTY } from "rxjs";
import { HTTP_INTERCEPTORS, HttpClientModule } from "@angular/common/http";
import { ErrorInterceptor } from "./core/interceptors/error.interceptor";
import { ApiInterceptor } from "./core/interceptors/api.interceptor";


@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    FooterComponent,
    HeaderComponent,
    AppRoutingModule,
    HttpClientModule,
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: ApiInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: ErrorInterceptor, multi: true },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
