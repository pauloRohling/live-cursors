import { ChangeDetectionStrategy, Component } from '@angular/core';

@Component({
  selector: 'app-canvas-page',
  standalone: true,
  imports: [],
  templateUrl: './canvas-page.component.html',
  styleUrl: './canvas-page.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class CanvasPageComponent {

}
